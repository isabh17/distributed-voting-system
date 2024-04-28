<template>
  <div id="app">
    <h1>Logs</h1>
    <table class="logs-table">
      <thead>
        <tr>
          <th>Banda</th>
          <th>Fecha y Hora</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item._id">
          <td>{{ item.key }}</td>
          <td>{{ item.Time }}</td>
        </tr>
      </tbody>
    </table>
    <button @click="refreshData">Update Data</button>
  </div>
</template>


