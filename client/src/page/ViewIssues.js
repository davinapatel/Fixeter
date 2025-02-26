import axios from "axios";
import {React, useState, useEffect} from 'react'
import { useParams} from 'react-router-dom';
import {Container, Card, Spinner, Button, Alert} from "react-bootstrap";
import { Link } from "react-router-dom";
import { format } from 'date-fns';

const ViewIssues = () => {

    const formatDate = (dateString) => {
        // Convert the string to a Date object
        const date = new Date(dateString);
    
        // Format the date (e.g., 4th January 2025, 12:04 PM)
        return format(date, "do MMMM yyyy, h:mm a");
    };

    const [apiData, setApiData] = useState([]);
    const [loading, setLoading] = useState(true);
    const { status } = useParams();


    const statusMap = {
        "logged": "Logged",
        "progress": "In Progress",
        "closed": "Closed"
    };

    const mappedStatus = statusMap[status] || "Unknown"

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

    const filteredData = Array.isArray(apiData)
        ? apiData.filter((record) => {
            return record.status === mappedStatus;
      })
    : [];

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
      <Container className="py-2">
        {filteredData.length === 0 ? (
                <>
                {[
                    'info',
                ].map((variant) => (
                    <Alert className="mt-3" key={variant} variant={variant}>
                        No items found for Status: {mappedStatus}
                    </Alert>
                ))}
                </>
            ) : (
            <div className="text-left">
                {filteredData && 
                filteredData.map((record, index) => (
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
            )}
            <Link to ="/staff-portal">
                <Button className="position-fixed bottom-0 end-0 m-3" variant="primary">Back to Portal</Button>
            </Link>
    </Container>
    );
};

export default ViewIssues;