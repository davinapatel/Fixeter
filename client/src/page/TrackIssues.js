import {React, useState, useEffect} from 'react';
import axios from "axios";
import {Container, Card, Spinner, Button} from "react-bootstrap";
import { Link } from "react-router-dom";
import { format } from 'date-fns';

const formatDate = (dateString) => {
    // Convert the string to a Date object
    const date = new Date(dateString);

    // Format the date (e.g., 4th January 2025, 12:04 PM)
    return format(date, "do MMMM yyyy, h:mm a");
};

const TrackIssues = () => {

    const [apiData, setApiData] = useState(false);
    const [loading, setLoading] = useState(true);

    

    useEffect(() => {
        const fetchData = async () => {
            try {
                const apiUrl = process.env.REACT_APP_API_ROOT;
                const response = await axios.get(apiUrl + "/issue");
                console.log("response", response)

                if (response.status === 200) {
                    if (response?.data.statusText === "Ok") {
                        setApiData(response?.data?.issue_records);
                    }
                }

                setLoading(false);
            } catch (error) {
                setLoading(false);
                console.log(error.response);
            }
        };

        fetchData();
        return () => {};
    }, [])

    console.log(apiData);

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
        <Container>
            <h3 className="mb-4"> Your Issues </h3>
            <div className="text-left">
                {apiData && 
                apiData.map((record, index) => (
                    <Card key={index} className="shadow-sm d-flex flex-row align-items-center" style={{ width: "60rem", borderRadius: "10px", overflow: "hidden" }}>
                        <Card.Img variant="left" src= {`http://localhost:8000/${record.image}`} className="img-fluid w-25 h-25"  style={{ width: "150px", height: "150px", objectFit: "cover" }} />
                        <Card.Body>
                            <Card.Title> {record.title} </Card.Title>
                            <Card.Text className="text-muted">{record.description}</Card.Text>
                            <Card.Text className="text-secondary small">Issue Published: {formatDate(record.date)}</Card.Text>
                        </Card.Body>
                        <div className="p-3">
                            <Button variant="primary" className='text-center'>Status: {record.status ? record.status: "Unknown"}</Button>
                        </div>
                    </Card>
                ))}               
            </div>
            <Link to ="/portal">
                <Button className="position-fixed bottom-0 end-0 m-3" variant="primary">Back to Portal</Button>
            </Link>
    </Container>
    );
};

export default TrackIssues;