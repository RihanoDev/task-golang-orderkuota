<template>
  <transition name="fade">
    <div v-if="message" :class="alertClass" class="alert">
      {{ message }}
    </div>
  </transition>
</template>

<script>
export default {
  props: {
    message: String,
    type: {
      type: String,
      default: "success",
    },
  },
  computed: {
    alertClass() {
      return this.type === "success" ? "alert alert-success" : "alert alert-danger";
    },
  },
  watch: {
    message() {
      if (this.message) {
        setTimeout(() => {
          this.$emit("clear-message");
        }, 3000);
      }
    },
  },
};
</script>

<style>
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}
</style>
