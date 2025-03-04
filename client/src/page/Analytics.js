import {React, useState, useEffect} from 'react';
import {Container} from "react-bootstrap";
import { Chart } from "react-google-charts";
import axios from "axios";

const Analytics = () => {

    const [apiData, setApiData] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const apiUrl = process.env.REACT_APP_API_ROOT;
                const response = await axios.get(apiUrl + "/issuehistory");
                console.log("response", response)

                if (response.status === 200) {
                    if (response?.data.statusText === "Ok") {
                        setApiData(response?.data?.issueHistory_records);
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

    const options = {
        title: "Count of Issues reported per Category in 2025"
    }

    const pieChartData = [
        ["Category", "Count"],
        ...apiData.map(item => [item.category, item.count])
    ];
        
    return (
        <Container className="py-2">
            <h3 className="text-center">Analytics of Issues Reported</h3>
            <Chart
                chartType="PieChart"
                data={pieChartData}
                options={options}
                wisth={"100%"}
                height={"400px"}
            />
        </Container>
    );
};

export default Analytics;