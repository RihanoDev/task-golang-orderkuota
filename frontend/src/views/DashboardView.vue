<template>
  <div class="container mt-4">
    <NavigationBar @logout="logout" />

    <!-- Sapaan User -->
    <h3 v-if="loggedInUser" class="mt-3">
      Hi, {{ loggedInUser.name || "User" }}! Selamat Datang!
    </h3>

    <!-- Statistik Pengguna -->
    <div class="mt-4 p-3 border rounded bg-light">
      <h5>Total Pengguna: {{ users.length }}</h5>
    </div>

    <!-- Notifikasi -->
    <Alert v-if="alertMessage" :message="alertMessage" :type="alertType" />

    <!-- Tabel Pengguna -->
    <UserTable :users="users" :loggedInUserId="loggedInUserId" @refresh="fetchUsers" />
  </div>
</template>

<script>
import NavigationBar from "@/components/NavigationBar.vue";
import UserTable from "@/components/UserTable.vue";
import Alert from "@/components/AlertMessage.vue";
import axios from "axios";

export default {
  components: { NavigationBar, UserTable, Alert },
  data() {
    return {
      users: [],
      loggedInUserId: sessionStorage.getItem("userId") || "",
      loggedInUser: null,
      alertMessage: "",
      alertType: "success",
    };
  },
  async created() {
    if (!sessionStorage.getItem("token") || !this.loggedInUserId) {
      console.warn("Token atau userId tidak ditemukan, logout otomatis.");
      this.logout();
      return;
    }
    await this.fetchUsers();
  },
  methods: {
    async fetchUsers() {
  const token = sessionStorage.getItem("token");
  if (!token) {
    console.warn("Token tidak ditemukan, logout otomatis.");
    this.logout();
    return;
  }

  try {
    console.log("Fetching users...");
    const response = await axios.get("http://localhost:9090/api/users", {
      headers: { Authorization: `Bearer ${token}` },
    });

    if (!response.data.data) {
      throw new Error("Data pengguna tidak ditemukan.");
    }

    this.users = response.data.data.map(user => ({
      id: String(user.ID),
      name: user.Name,
      email: user.Email,
      created_at: user.CreatedAt,
      updated_at: user.UpdatedAt
    }));

    console.log("User ID dari sessionStorage:", this.loggedInUserId);
    console.log("List users dari API:", this.users);

    this.loggedInUser = this.users.find(user => String(user.id).trim() === String(this.loggedInUserId).trim());

    if (!this.loggedInUser) {
      console.warn("User tidak ditemukan dalam daftar, logout otomatis.");
      this.logout();
      return;
    }

    console.log("User berhasil ditemukan:", this.loggedInUser);
  } catch (error) {
    console.error("Gagal mengambil data pengguna", error);

    if (error.response && error.response.status === 401) {
      console.warn("Token kadaluarsa, logout otomatis.");
      this.logout();
    } else {
      this.alertMessage = "Gagal memuat data pengguna!";
      this.alertType = "danger";
    }
  }
},
logout() {
  sessionStorage.clear();
  console.log("🚪 Logout berhasil, redirect ke login...");
  this.$router.push("/login");
},
  }
};
</script>
