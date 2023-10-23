import React from 'react';

function NetworkSlice({ name, bandwidth, latency, serviceType, timeOfDay }) {
  return (
    <div className="network-slice">
      <h2>{name}</h2>
      <p>Bandwidth: {bandwidth} Mbps</p>
      <p>Latency: {latency} ms</p>
      <p>Service Type: {serviceType}</p>
      <p>Time of Day: {timeOfDay}</p>
    </div>
  );
}

export default NetworkSlice;
