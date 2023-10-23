import React, { useState, useEffect } from 'react';
import NetworkSlice from './NetworkSlice';
import EfficiencyTable from './EfficiencyTable';

function Dashboard() {
  const [efficiencyData, setEfficiencyData] = useState([]);

  useEffect(() => {
    fetch('http://your-go-backend-url/api/efficiency')
      .then((response) => response.json())
      .then((data) => setEfficiencyData(data));
  }, []);

  return (
    <div>
     <NetworkSlice />
     <EfficiencyTable />
    </div>
  );
}

export default Dashboard;
