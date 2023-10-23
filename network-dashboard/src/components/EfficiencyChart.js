import React, { useEffect, useRef } from 'react';
import Chart from 'chart.js/auto';

function EfficiencyChart({ efficiencyData }) {
  const chartRef = useRef();

  useEffect(() => {
    if (efficiencyData.length === 0) {
      return; // Don't create the chart if there's no data
    }

    // Extract labels and values from the efficiencyData
    const labels = efficiencyData.map((item) => item.TimeOfDay);
    const values = efficiencyData.map((item) => item.Efficiency);

    // Create the Chart.js chart
    const ctx = chartRef.current.getContext('2d');
    new Chart(ctx, {
      type: 'bar',
      data: {
        labels: labels,
        datasets: [
          {
            label: 'Efficiency',
            data: values,
            backgroundColor: 'rgba(75, 192, 192, 0.2)',
            borderColor: 'rgba(75, 192, 192, 1)',
            borderWidth: 1,
          },
        ],
      },
    });
  }, [efficiencyData]);

  return <canvas ref={chartRef} />;
}

export default EfficiencyChart;
