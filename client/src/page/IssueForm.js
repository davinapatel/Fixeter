import axios from "axios";
import React, { useState } from "react";
import { Col, Container, Row, Spinner } from "react-bootstrap";

import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";

const Add = () => {
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const saveForm = async (data) => {
    setLoading(true);
    // console.log(data);

    // data.file = data.image[0];
    // data.image = null;

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
            {/* <Col xs="12" className="py-3">
              <label>Image</label>
              <input
                type="text"
                defaultValue=""
                className={`${errors.image && "error"}`}
                placeholder="Placeholder for an image upload button"
                {...register("image", {
                  required: {
                    value: true,
                  },
                })}
              />
            </Col> */}
            {/* <Col xs="12" className="py-3">
              <label>Image</label>
              <input
                type="file"
                className={`${errors.image && "error"}`}
                placeholder="Please enter content"
                {...register("image")}
              />
            </Col> */}
            <Col>
              <button type="submit">Submit Issue</button>
            </Col>
          </Row>
        </form>
      </Container>
    </>
  );
};

export default Add;