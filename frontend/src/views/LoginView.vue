<template>
  <div class="container d-flex justify-content-center align-items-center vh-100">
    <div class="card p-4 shadow-lg rounded" style="width: 400px">
      <h2 class="text-center mb-3">Login</h2>

      <AlertMessage v-if="errorMessage" :message="errorMessage" type="danger" />

      <form @submit.prevent="login" class="d-flex flex-column gap-3">
        <FormInput label="Email" type="email" v-model="email" id="email" required />
        <FormInput label="Password" type="password" v-model="password" id="password" required />

        <button type="submit" class="btn btn-primary w-100">Login</button>
      </form>

      <p class="text-center mt-3">
        Belum punya akun?
        <router-link to="/register" class="text-primary text-decoration-none">Daftar di sini</router-link>
      </p>
    </div>
  </div>
</template>

<script>
import AlertMessage from "@/components/AlertMessage.vue";
import FormInput from "@/components/FormInput.vue";
import axios from "axios";

export default {
  components: { AlertMessage, FormInput },
  data() {
    return {
      email: "",
      password: "",
      errorMessage: "",
    };
  },
  methods: {
  async login() {
    console.log("Login function dipanggil");

    try {
      const response = await axios.post("http://localhost:9090/login", {
        email: this.email,
        password: this.password,
      });

      console.log("Response dari backend:", response.data);
      const token = response.data.data.access_token;
      const userId = response.data.data.user.id;

      if (!token || !userId) {
        throw new Error("Token tidak ditemukan || Id tidak ditemukan");
      }

      sessionStorage.setItem("token", token);
      sessionStorage.setItem("userId", userId);

      console.log("Token tersimpan:", sessionStorage.getItem("token"));
      console.log("User ID tersimpan:", sessionStorage.getItem("userId"));

      this.$router.push("/dashboard");
    } catch (error) {
       console.error("Login gagal:", error);

        if (error.response) {
          this.errorMessage = error.response.data.message || "Login gagal, periksa email dan password!";
        } else {
          this.errorMessage = "Terjadi kesalahan, coba lagi nanti!";
        }
    }
  },
},
};
</script>
