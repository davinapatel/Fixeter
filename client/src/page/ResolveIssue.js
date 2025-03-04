import axios from "axios";
import {React, useState, useEffect} from 'react';
import {Container, Row, Spinner, Form, Image, Button} from "react-bootstrap";
import { useForm } from "react-hook-form";
import {useParams, Link, useNavigate} from "react-router-dom";
import { format } from 'date-fns';
import Map, { Marker, NavigationControl} from "react-map-gl/mapbox";
import 'mapbox-gl/dist/mapbox-gl.css';

const formatDate = (dateString) => {
    // Convert the string to a Date object
    const date = new Date(dateString);

    // Format the date (e.g., 4th January 2025, 12:04 PM)
    return format(date, "do MMMM yyyy, h:mm a");
};

const ResolveIssue = () => {
    const {
        register,
        handleSubmit,
        setValue,
        formState: { errors },
      } = useForm();

    const {status, recordId} = useParams();
    const [loading, setLoading] = useState(true);
    const [apiData, setApiData] = useState(null);
    const [specificResource, setSpecificResource] = useState(null);
    const [resources, setResources] = useState(null);
    const [checked, setChecked] = useState(false);
    const navigate = useNavigate();

    const saveForm = async (data) => {
        setLoading(true);
        if (status === "logged"){
            apiData.resourceId = Number(data.resource)
            apiData.status = "In Progress"
        };
        if (status === "progress") {
            apiData.comments = data.comments
            if (checked === true) {
                apiData.status = "Closed"
            };
        }
               
        try {
            const apiUrl = process.env.REACT_APP_API_ROOT;
            const response = await axios.put(apiUrl + "/issue/" + recordId, apiData, {
                headers: {
                    "Content-Type": "multipart/form-data",
                },  
            });
             if (response.status === 200) {
                console.log(response);
                navigate("/staff-portal")
             }
             setLoading(false); 
        } catch (error) {
            setLoading(false);
            console.log(error.response);
        }
    };

    const handleCheckbox = (event) => {
        setChecked(event.target.checked);

    }

    useEffect(() => {
        const fetchData = async () => {
            try {
              const apiUrl = process.env.REACT_APP_API_ROOT;
              const response = await axios.get(apiUrl + "/issue/" + recordId);

              if (response.status === 200) {
                   if (response?.data.statusText === "Ok") {
                        setApiData(response?.data?.record)
                    }
                };
                

            } catch (error) {
                console.log(error.response);
            } finally {
                setLoading(false);
            }

        };

        fetchData();
    }, [recordId])

    useEffect (() => {
        const fetchSpecificResource = async () => {
            try {
                if (status === "progress" && apiData?.resourceId) {
                    const apiUrl = process.env.REACT_APP_API_ROOT;
                    const resourceResponse = await axios.get(apiUrl + "/resource/" + apiData.resourceId)
                    setSpecificResource(resourceResponse?.data?.record)
                    }
                else {
                    console.log("Not needed")
                }
            
            } catch (error) {
                console.log(error.response);
            } finally {
                setLoading(false);
            }
        };
        fetchSpecificResource();
    }, [apiData, status])

    useEffect(() => {
        const fetchResources = async () => {
            try {
                const apiUrl = process.env.REACT_APP_API_ROOT;
                const response2 = await axios.get(apiUrl + "/resource");

                if (response2.status === 200) {
                    if (response2?.data?.statusText === "Ok") {
                        setResources(response2?.data?.resource_records);
                    }
                }
                setLoading(false);
            } catch (error) {
                console.log("Error fetching resources:", error.response);
            } finally {
                setLoading(false);
            }
        };

        fetchResources();
    }, []);

    useEffect(() => {
        if (apiData?.comments) {
            setValue("comments", apiData.comments); // Set default value when data loads
        }
    }, [apiData, setValue]);

    if (loading || !apiData || !resources || (status === "progress" && !specificResource)) {
        return (
            <>
                <Container className="spinner">
                    <Spinner animation="grow" />
                </Container>
            </>
        );
    }

    const latitude = Number(apiData?.latitude);
    const longitude = Number(apiData?.longitude);

    

    

    return (
        <Container className="py-2">
        <Row>
          <h3>
            Resolve Issue
          </h3>
          <div style={{ display: "flex", gap: "20px", alignItems: "flex-start", flexWrap: "wrap" }}>
            <div style={{ flex: 1, minWidth: "300px" }}>
                <Form onSubmit={handleSubmit(saveForm)}>
                    <Form.Group className="mb-3">
                        <Form.Label className="left-align">Category</Form.Label>
                        <Form.Select disabled>
                            <option>{apiData.category}</option>
                        </Form.Select>
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label className="left-align">Title of Issue</Form.Label>
                        <Form.Control placeholder={apiData.title} disabled />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label className="left-align">Date of Issue</Form.Label>
                        <Form.Control placeholder={formatDate(apiData.date)} disabled />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label className="left-align">Description</Form.Label>
                        <Form.Control placeholder={apiData.description} disabled />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label className="left-align">Address</Form.Label>
                        <Form.Control placeholder={apiData.address} disabled />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label className="left-align">Allocate Resource</Form.Label>
                            {status === "logged" && (
                            <Form.Select className={`${errors.resource && "error"}`}
                                {...register("resource", {
                                required: { value: true, message: "Resource is required." },
                                })}
                            >
                                { resources?.map((resource, index) => (
                                    <option key={resource.id} value={resource.id} disabled={index === 0}>
                                        {resource.type}
                                    </option>
                                ))}
                            </Form.Select>
                            )}

                            {status === "progress" && (
                            <Form.Control placeholder={specificResource.type} disabled />
                            )}
                        
                        {errors.resource && (
                        <div className="error text-danger">{errors.resource.message}</div>
                      )}
                    </Form.Group>
                    {status === "progress" && (
                        <Form.Group className="mb-3">
                            <Form.Label className="left-align">Department to Resolve</Form.Label>
                            <Form.Control placeholder={specificResource.department} disabled />
                            <Form.Label className="left-align, mt-3">Comments</Form.Label>
                            <Form.Control 
                                as="textarea"
                                rows={3}
                                // value= {apiData.comments}
                                placeholder = "Please provide an update"
                                {...register("comments", { required: "Comments are required." })}
                            />                        
                            <Form.Check
                                type="checkbox"
                                label = "Close Issue"
                                className="mt-3"
                                checked = {checked}
                                onChange={handleCheckbox}>

                            </Form.Check>
                        </Form.Group>
                            
                        
                      )}
                
                  <Button type="submit">Save</Button>
                  <Link to="/staff-portal">
                    <Button>Back to Staff Portal</Button>
                  </Link>
                </Form>
            </div>

            <div style={{ display: "flex", flexDirection: "column", gap: "10px", minWidth: "300px" }}>
                <div style={{ width: "350px", height: "250px", borderRadius: "8px", overflow: "hidden" }}>
                    <Map
                        initialViewState={{
                            latitude: latitude,
                            longitude: longitude,
                            zoom: 15,
                        }}
                        mapboxAccessToken={process.env.REACT_APP_MAPBOX_TOKEN}
                        mapStyle= "mapbox://styles/mapbox/streets-v11"
                        style={{ width: "100%", height: "100%" }}
                    >
                        <Marker
                            latitude= {latitude}
                            longitude={longitude}
                            color="red"
                        />
                        <NavigationControl position="bottom-right" />
                    </Map>
                </div>

                <div style={{ width: "350px", height: "250px", borderRadius: "8px", overflow: "hidden" }}>
                    <Image src={`http://localhost:8000/${apiData.image}`} fluid style={{ width: "100%", height: "100%", objectFit: "cover" }} />
                </div>
            </div>
        </div>
        </Row>
      </Container>
    );
};

export default ResolveIssue;