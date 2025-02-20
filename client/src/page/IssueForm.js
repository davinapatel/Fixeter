import axios from "axios";
import React, { useState } from "react";
import { Col, Container, Nav, Row, Spinner } from "react-bootstrap";
import Map, { Marker, NavigationControl, Room } from "react-map-gl/mapbox";
import 'mapbox-gl/dist/mapbox-gl.css';

import { set, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";

const Add = () => {
  const [loading, setLoading] = useState(false);

  const [marker, setMarker] = useState({
    latitude: 50.7256138,
    longitude: -3.5269209,
  });

  const [formData, setFormData] = useState({
    latitude: 50.7256138,
    longitude: -3.5269209,
    address: "",
  });

  const [address, setAddress] = useState({
    address: ""
  });

  const fetchAddressFromCoords = async (lng, lat) => {
    try {
      console.log("https://api.mapbox.com/geocoding/v5/mapbox.places/" + lng + "," + lat + ".json")
      const resp = await axios.get(
        "https://api.mapbox.com/geocoding/v5/mapbox.places/" + lng + "," + lat + ".json",
        {
          params: {
            access_token: process.env.REACT_APP_MAPBOX_TOKEN,
          },
        }
      );

      if (resp.data.features.length > 0) {
        const address = resp.data.features[0].place_name;
        setAddress(() => ({
          address: address,
        }));
      }
    } catch (error) {
      console.error("Error fetching address:", error);
    }
  };

  const handleMouseClick = (event) => {
    console.log("Map click event:", event);
    console.log(event.lngLat);
    fetchAddressFromCoords(event.lngLat.lng, event.lngLat.lat);
    setMarker({
      latitude: event.lngLat.lat,
      longitude: event.lngLat.lng,
    });

    setFormData(() => ({
      latitude: event.lngLat.lat,
      longitude: event.lngLat.lng,
    }));

    
  };

  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const saveForm = async (data) => {
    setLoading(true);
    try {
      const apiUrl = process.env.REACT_APP_API_ROOT;
      const response = await axios.post(apiUrl + "/issue", data, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
        body: JSON.stringify({ title: String(data.title) }),
      });

      if (response.status === 201) {
        console.log(response);
        navigate("/");
      }

      setLoading(false);
    } catch (error) {
      setLoading(false);
      console.log(error.response);
    }
  };

  if (loading) {
    return (
      <>
        <Container className="spinner">
          <Spinner animation="grow" />
        </Container>
      </>
    );
  }

  return (
    <>
      <Container>
        <div class="parent">
          <div class="child">
            <h1>Log a New Issue</h1>
            <form onSubmit={handleSubmit(saveForm)}>
              <Row>
                <Col xs="12" className="py-3">
                  <label>Category</label>
                  <input
                    defaultValue=""
                    className={`${errors.category && "error"}`}
                    placeholder="Please select a Category"
                    {...register("category", {
                      required: { value: true, message: "Category is required." },
                    })}
                  />
                  {errors.category && (
                    <div className="error">{errors.category.message}</div>
                  )}
                </Col>
                <Col xs="12" className="py-3">
                  <label>Title of Issue</label>
                  <input
                    defaultValue=""
                    className={`${errors.title && "error"}`}
                    placeholder="Please enter a Title"
                    {...register("title", {
                      required: {
                        value: true,
                        message: "Title is required.",
                      },
                    })}
                  />
                  {errors.title && (
                    <div className="error">{errors.title.message}</div>
                  )}
                </Col>
                <Col xs="12" className="py-3">
                  <label>Date of Issue</label>
                  <input
                    defaultValue=""
                    type="datetime-local"
                    className={`${errors.date && "error"}`}
                    placeholder="Please enter a date"
                    {...register("date", {
                      required: {
                        value: true,
                        message: "Date is required.",
                      },
                    })}
                  />
                  {errors.date && (
                    <div className="error">{errors.date.message}</div>
                  )}
                </Col>
                <Col xs="12" className="py-3">
                  <label>Description</label>
                  <input
                    type="text"
                    defaultValue=""
                    className={`${errors.title && "error"}`}
                    placeholder="Please enter a Description"
                    {...register("description", {
                      required: {
                        value: true,
                        message: "Description is required.",
                      },
                    })}
                  />
                  {errors.description && (
                    <div className="error">{errors.description.message}</div>
                  )}
                </Col>
                <Col xs="12" className="py-3">
                  <label>Latitude</label>
                  <input
                    type="float64"
                    defaultValue=""
                    value={formData.latitude}
                    className={`${errors.latitude && "error"}`}
                    placeholder="Please select a location on Map"
                    readOnly
                    {...register("latitude", {
                      required: {
                        value: true,
                        message: "Latitude is required via Map Selection.",
                      },
                    })}
                  />
                  {errors.latitude && (
                    <div className="error">{errors.latitude.message}</div>
                  )}
                </Col>
                <Col xs="12" className="py-3">
                  <label>Longitude</label>
                  <input
                    type="float64"
                    defaultValue=""
                    value={formData.longitude}
                    className={`${errors.longitude && "error"}`}
                    placeholder="Please select a location on Map"
                    readOnly
                    {...register("longitude", {
                      required: {
                        value: true,
                        message: "Longitude is required via Map Selection.",
                      },
                    })}
                  />
                  {errors.longitude && (
                    <div className="error">{errors.longitude.message}</div>
                  )}
                </Col>
                <Col xs="12" className="py-3">
                  <label>Address</label>
                  <input
                    type="text"
                    defaultValue=""
                    value={address.address}
                    className={`${errors.address && "error"}`}
                    placeholder="Please select a location on Map"
                    readOnly
                    {...register("address", {
                      required: {
                        value: true,
                        message: "Address is required via Map Selection.",
                      },
                    })}
                  />
                  {errors.address && (
                    <div className="error">{errors.address.message}</div>
                  )}
                </Col>               
                <Col>
                  <button type="submit">Submit Issue</button>
                </Col>
              </Row>
            </form>
          </div>
          <div class="child" style={{ width: "50vw", height: "50vh", marginTop: "20px" }}>
            <Map
              initialViewState={{
                latitude: 50.7256138,
                longitude: -3.5269209,
                zoom: 10,
              }}
              mapboxAccessToken={process.env.REACT_APP_MAPBOX_TOKEN}
              mapStyle= "mapbox://styles/mapbox/streets-v11"
              style={{ width: "100%", height: "100%" }}
              onDblClick={handleMouseClick}
            >
                <Marker
                  latitude={marker.latitude}
                  longitude={marker.longitude}
                  color="red"
                />
                <NavigationControl position="bottom-right" />
            </Map>
        </div>
       </div>
      </Container>
    </>
  );
};

export default Add;