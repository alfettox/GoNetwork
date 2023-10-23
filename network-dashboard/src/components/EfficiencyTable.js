import React, { useState, useEffect } from 'react';

function EfficiencyTable() {
  const [data, setData] = useState([]);

  useEffect(() => {
    // Fetch data from the Go API endpoint
    fetch('http://localhost:8080/api/efficiency') 
    .then((response) => response.json())
    .then((data) => {
      setData(data);
    })
    .catch((error) => {
      console.error('Error fetching data:', error);
    });
}, []);


  return (
    <div>
      <h1>Efficiency Data</h1>
      <table>
        <thead>
          <tr>
            <th>Service Type</th>
            <th>Time of Day</th>
            <th>Efficiency</th>
          </tr>
        </thead>
        <tbody>
          {data.map((item, index) => (
            <tr key={index}>
              <td>{item.ServiceType}</td>
              <td>{item.TimeOfDay}</td>
              <td>{item.Efficiency}%</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default EfficiencyTable;
