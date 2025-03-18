<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container-fluid">
      <router-link class="navbar-brand" to="/dashboard">Dashboard</router-link>
      <button class="btn btn-outline-light" @click="handleLogout">Logout</button>
    </div>
  </nav>
</template>

<script>
export default {
  methods: {
    handleLogout() {
      this.$emit("logout");
    },
  },
};
</script>
