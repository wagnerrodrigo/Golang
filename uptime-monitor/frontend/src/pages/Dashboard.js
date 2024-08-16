import React, { useState, useEffect } from 'react';
import UptimeTable from '../components/UptimeTable';
import { fetchUptimeData } from '../services/api';

function Dashboard() {
    const [data, setData] = useState([]);

    useEffect(() => {
        fetchUptimeData()
            .then(response => {
                console.log("Fetched data:", response);  // Adiciona log
                setData(response);
            })
            .catch(error => {
                console.error("Error fetching data:", error);
            });
    }, []);

    return (
        <div>
            <h1>UpTime Monitor</h1>
            <UptimeTable data={data} />
        </div>
    );
}


export default Dashboard;
