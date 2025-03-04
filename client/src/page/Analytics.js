import {React, useState, useEffect} from 'react';
import {Container, Spinner} from "react-bootstrap";
import { Chart } from "react-google-charts";
import axios from "axios";

const Analytics = () => {

    const [apiData, setApiData] = useState([]);
    const [chartApiData, setChartApiData] = useState([]);
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

    useEffect(() => {
        const fetchChartData = async () => {
            try {
                const apiUrl = process.env.REACT_APP_API_ROOT;
                const response2 = await axios.get(apiUrl + "/issueStatus");
                console.log("response2", response2)

                if (response2.status === 200 ) {
                    if (response2?.data.statusText === "Ok") {
                        console.log('in here')
                        setChartApiData(response2?.data?.issueStatus_records);
                    }
                }

                setLoading(false);
            } catch (error) {
                setLoading(false);
                console.log(error.response);
            }
        };
        fetchChartData();
        return () => {};
    }, [])


    const options = {
        title: "Count of Issues reported per Category in 2025"
    }

    const pieChartData = [
        ["Category", "Count"],
        ...apiData.map(item => [item.category, item.count])
    ];

    const chartData = [
        ["Status", "Count"],
        ...chartApiData.map(x => [x.Status, Number(x.Count)])
    ];

    const columnChartOptions = {
        title: "Current Total of Issues by Status",
        chartArea: { width: "70%" },
        hAxis: {
          title: "Status",
        },
        vAxis: {
          title: "Count",
          minValue: 0,
        },
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
    console.log(chartData)
    return (
        <Container className="py-2">
            <h3 className="text-center">Analytics of Issues Reported</h3>
            <div>
            <Chart
                chartType="PieChart"
                data={pieChartData}
                options={options}
                wisth={"100%"}
                height={"400px"}
            />
            <Chart 
                chartType="ColumnChart"
                width="100%"
                height="100%"
                options = {columnChartOptions}
                data={chartData} />
            </div>
        </Container>
    );
};

export default Analytics;