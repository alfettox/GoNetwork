const data = {
    labels: ["Category 1", "Category 2", "Category 3", "Category 4"],
    datasets: [
      {
        label: "My Dataset",
        data: [10, 20, 30, 40],

        xAxisID: 'x-axis-0',
        fill: false,
        borderColor: 'rgb(75, 192, 192)',
        tension: 0.1,
      },
    ],
  };
  
  const config = {
    type: 'line',
    data: data,
    options: {
      scales: {
        x: {
          type: 'category',
          title: {
            display: true,
            text: 'X-Axis Label',
          },
        },
        y: {
          beginAtZero: true,
          title: {
            display: true,
            text: 'Y-Axis Label',
          },
        },
      },
    },
  };
  
  const myChart = new Chart(document.getElementById('myChart'), config);
  