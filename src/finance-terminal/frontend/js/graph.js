const xlables = [];
const assetsByMonth = [];
const debtsByMonth = [];

createChart();
async function getData(){
    const response = await fetch('./finances.csv');
    const data = await response.text();
    const rows = data.split('\n')
    rows.pop()
    rows.forEach(elt => {
        const row = elt.split(',');
        const date = row[0]
        xlables.push(date)
        const assets = row[1]
        assetsByMonth.push(assets);
        const debts = row[2]
        debtsByMonth.push(debts);
        console.log(assets);
    });
}
async function createChart(){
   await getData();
   const ctx = document.getElementById('myChart').getContext('2d');
   const myChart = new Chart(ctx, {
       type: 'line',
       data: {
           labels: xlables,
           datasets: [{
               label: 'Assets',
               data: assetsByMonth,
               backgroundColor: [
                   'rgba(72, 143, 49, .75)',
               ],
               borderColor: [
                   'rgba(72, 143, 49, .75)',
               ],
               borderWidth: 1
           },{
               label: 'Debts',
               data: debtsByMonth,
               backgroundColor: [
                   'rgba(232, 48, 48, .5)',
               ],
               borderColor: [
                   'rgba(232, 48, 48, .5)',
               ],
               borderWidth: 1
           },
       ]
       },
       options: {
             scales: {
                 yAxes: [{
                   ticks: {
                   beginAtZero: true
                   }
                 }]
         }
       }
   });

 }