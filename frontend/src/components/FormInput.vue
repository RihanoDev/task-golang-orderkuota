<template>
  <div class="mb-3">
    <label :for="id" class="form-label">{{ label }}</label>
    <input
      :type="type"
      :value="modelValue"
      @input="$emit('update:modelValue', $event.target.value)"
      class="form-control"
      :id="id"
      required
    />
  </div>
</template>

<script>
export default {
  props: {
    modelValue: String,
    type: String,
    id: String,
    label: String
  }
}
</script>