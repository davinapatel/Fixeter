import axios from "axios";
import {React, useState, useEffect} from 'react';
import {Container, Row, Spinner, Form, Image, Button} from "react-bootstrap";
import { useForm } from "react-hook-form";
import {useParams, Link} from "react-router-dom";
import { format } from 'date-fns';
import Map, { Marker, NavigationControl} from "react-map-gl/mapbox";
import 'mapbox-gl/dist/mapbox-gl.css';

const formatDate = (dateString) => {
    // Convert the string to a Date object
    const date = new Date(dateString);

    // Format the date (e.g., 4th January 2025, 12:04 PM)
    return format(date, "do MMMM yyyy, h:mm a");
};

const Home = () => {
    const {
        register,
        handleSubmit,
        formState: { errors },
      } = useForm();

    const {recordId} = useParams();
    const [loading, setLoading] = useState(true);
    const [apiData, setApiData] = useState(null);
    const [resources, setResources] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const apiUrl = process.env.REACT_APP_API_ROOT;
                const response = await axios.get(apiUrl + "/issue/" + recordId);

                if (response.status === 200) {
                    if (response?.data.statusText === "Ok") {
                        setApiData(response?.data?.record);
                    }
                }

            } catch (error) {
                console.log(error.response);
            } finally {
                setLoading(false);
            }

        };

        fetchData();
    }, [recordId])

    useEffect(() => {
        const fetchResources = async () => {
            try {
                const apiUrl = process.env.REACT_APP_API_ROOT;
                const response2 = await axios.get(apiUrl + "/resource");

                if (response2.status === 200) {
                    if (response2?.data?.statusText === "Ok") {
                        setResources(response2?.data?.resource_records);
                    }
                    console.log(response2)
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

    if (loading) {
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
                <Form>
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
                        <Form.Select className={`${errors.resource ? "error" : ""}`}
                            {...register("resource", {
                            required: { value: true, message: "Resource is required." },
                            })}
                            defaultValue=""
                        >
                            <option value="" disabled >Please select a resource</option>
                            <option >{resources[1].type}</option>
                            <option>{resources[2].type}</option>
                            <option>{resources[3].type}</option>
                            <option>{resources[4].type}</option>
                        </Form.Select>
                        {errors.resource && (
                        <div className="error text-danger">{errors.resource.message}</div>
                      )}
                    </Form.Group>
                </Form>
                <div className="d-flex gap-2">
                    <Button type="submit">Save</Button>
                    <Link to="/staff-portal">
                        <Button>Back to Staff Portal</Button>
                    </Link>
                </div>
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

export default Home