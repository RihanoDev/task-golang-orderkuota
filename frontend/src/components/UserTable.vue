<template>
  <div>
    <table class="table table-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Nama</th>
          <th>Email</th>
          <th>Dibuat Pada</th>
          <th>Diubah Pada</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td>{{ user.id }}</td>
          <td>{{ user.name }}</td>
          <td>{{ user.email }}</td>
          <td>{{ formatDate(user.created_at) }}</td>
          <td>{{ formatDate(user.updated_at) }}</td>
          <td>
            <button class="btn btn-sm btn-warning me-2" @click="editUser(user)">Edit</button>
            <button
              class="btn btn-sm btn-danger"
              @click="deleteUser(user.id)"
              :disabled="String(user.id).trim() === String(loggedInUserId).trim()"
            >
              Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>

    <!-- Modal Edit User -->
    <div v-if="showEditModal" class="modal d-block">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Edit User</h5>
            <button type="button" class="btn-close" @click="showEditModal = false"></button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="updateUser">
              <div class="mb-3">
                <label for="name" class="form-label">Nama</label>
                <input type="text" id="name" class="form-control" v-model="editedUser.name" required />
              </div>
              <div class="mb-3">
                <label for="email" class="form-label">Email</label>
                <input type="email" id="email" class="form-control" v-model="editedUser.email" required />
              </div>
              <div class="mb-3">
                <label for="password" class="form-label">Password (Opsional)</label>
                <input type="password" id="password" class="form-control" v-model="editedUser.password" placeholder="Biarkan kosong jika tidak ingin mengubah" />
              </div>
              <div class="modal-footer">
                <button type="button" class="btn btn-secondary" @click="showEditModal = false">Batal</button>
                <button type="submit" class="btn btn-primary">Simpan</button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  props: {
    users: Array,
    loggedInUserId: [String, Number],
  },
  data() {
    return {
      showEditModal: false,
      editedUser: { id: "", name: "", email: "", password: "" },
    };
  },
  methods: {
    formatDate(dateString) {
      if (!dateString) return "-";
      const date = new Date(dateString);
      return isNaN(date) ? "-" : date.toLocaleString();
    },
    editUser(user) {
      this.editedUser = { ...user, password: "" };
      this.showEditModal = true;
    },
    async updateUser() {
      try {
        const token = sessionStorage.getItem("token");
        const updateData = { ...this.editedUser };

        if (!updateData.password) delete updateData.password;

        await axios.put(`http://localhost:9090/api/users/${this.editedUser.id}`, updateData, {
          headers: { Authorization: `Bearer ${token}` },
        });

        this.$emit("show-alert", "User berhasil diperbarui!", "success");
        this.$emit("refresh");
        this.showEditModal = false;
      } catch (error) {
        console.error("Gagal memperbarui user", error);
        this.$emit("show-alert", "Gagal memperbarui user!", "danger");
      }
    },
    async deleteUser(id) {
  if (!confirm("Apakah Anda yakin ingin menghapus user ini?")) return;

  try {
    const token = sessionStorage.getItem("token");
    if (!token) {
      console.warn("Token tidak ditemukan, logout otomatis.");
      window.location.replace("/login");
      return;
    }

    await axios.delete(`http://localhost:9090/api/users/${id}`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    if (String(id).trim() === String(this.loggedInUserId).trim()) {
      console.warn("User menghapus akunnya sendiri, logout otomatis.");
      sessionStorage.removeItem("token");
      sessionStorage.removeItem("userId");
      alert("User berhasil dihapus! Anda akan diarahkan ke halaman login.");
      window.location.replace("/login");
    } else {
      this.$emit("show-alert", "User berhasil dihapus!", "success");
      this.$emit("refresh");
    }
  } catch (error) {
    console.error("Gagal menghapus user", error);
    this.$emit("show-alert", "Gagal menghapus user!", "danger");
  }
}
  },
};
</script>
