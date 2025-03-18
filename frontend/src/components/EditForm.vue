<template>
  <div v-if="show" class="modal d-block">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">Edit User</h5>
          <button type="button" class="btn-close" @click="closeModal"></button>
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
              <button type="button" class="btn btn-secondary" @click="closeModal">Batal</button>
              <button type="submit" class="btn btn-primary">Simpan</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  props: {
    user: Object,
    show: Boolean,
  },
  data() {
    return {
      editedUser: { ...this.user, password: "" },
    };
  },
  watch: {
    user(newUser) {
      this.editedUser = { ...newUser, password: "" };
    },
  },
  methods: {
    async updateUser() {
      try {
        const token = sessionStorage.getItem("token");
        const updateData = { ...this.editedUser };

        if (!updateData.password) delete updateData.password;

        await axios.put(`http://localhost:9090/api/users/${this.editedUser.id}`, updateData, {
          headers: { Authorization: `Bearer ${token}` },
        });

        alert("User berhasil diperbarui!");
        this.$emit("updated");
        this.closeModal();
      } catch (error) {
        console.error("Gagal memperbarui user", error);
        alert("Gagal memperbarui user!");
      }
    },
    closeModal() {
      this.$emit("close");
    },
  },
};
</script>

<style scoped>
.modal {
  background: rgba(0, 0, 0, 0.5);
}
</style>
