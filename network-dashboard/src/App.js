import React, { useState, useEffect } from 'react';
import EfficiencyTable from './components/EfficiencyTable'; 
import EfficiencyChart from './components/EfficiencyChart';

function App() {
  const [efficiencyData, setEfficiencyData] = useState([]);

  useEffect(() => {
    // Fetch data from the Go API endpoint
    fetch('http://localhost:8080/api/efficiency')
      .then((response) => response.json())
      .then((data) => {
        setEfficiencyData(data);
      })
      .catch((error) => {
        console.error('Error fetching data:', error);
      });
  }, []);

  return (
    <div>
      <EfficiencyTable data={efficiencyData} />
      <EfficiencyChart efficiencyData={efficiencyData} />
    </div>
  );
}

export default App;
