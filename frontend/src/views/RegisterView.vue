<template>
  <div class="container d-flex justify-content-center align-items-center vh-100">
    <div class="card p-4 shadow-lg" style="width: 400px">
      <h2 class="text-center">Register</h2>

      <AlertMessage v-if="errorMessage" :message="errorMessage" type="danger" />
      <AlertMessage v-if="successMessage" :message="successMessage" type="success" />

      <form @submit.prevent="register">
        <FormInput label="Nama" type="text" v-model="name" id="name" />
        <FormInput label="Email" type="email" v-model="email" id="email" />
        <FormInput label="Password" type="password" v-model="password" id="password" />

        <button type="submit" class="btn btn-success w-100 mt-2">Register</button>
      </form>

      <p class="text-center mt-3">
        Sudah punya akun?
        <router-link to="/login">Login di sini</router-link>
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
      name: "",
      email: "",
      password: "",
      errorMessage: "",
      successMessage: "",
    };
  },
  methods: {
    async register() {
      try {
        await axios.post("http://localhost:9090/register", {
          name: this.name,
          email: this.email,
          password: this.password,
        });

        this.successMessage = "Registrasi berhasil! Silakan login.";
        this.errorMessage = "";
      } catch (error) {
        this.errorMessage = "Registrasi gagal!";
        this.successMessage = "";
      }
    },
  },
};
</script>