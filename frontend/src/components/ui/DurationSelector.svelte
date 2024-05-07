<script lang="ts">
  export let value = "";
  export let placeholder = "";
  export let required = false;

  let isValid = true;

  let rgx =
    /^([0-9]+(\.[0-9]+)?h)?([0-9]+(\.[0-9]+)?m)?([0-9]+(\.[0-9]+)?s)?([0-9]+(\.[0-9]+)?ms)?([0-9]+(\.[0-9]+)?[uµ]s)?([0-9]+(\.[0-9]+)?ns)?$/gm;

  $: {
    if (value) {
      isValid = rgx.test(value);
    } else {
      isValid = !required;
    }
  }
</script>

<div
  class="wrapper w-full h-10 border overflow-hidden border-white/20 rounded bg-default-bg flex justify-center items-center"
  class:border-red-500={!isValid}
>
  <input
    class="bg-default-bg w-full h-full px-3 outline-none"
    class:text-red-500={!isValid}
    type="text"
    bind:value
    {placeholder}
    autocorrect="off"
    autocapitalize="none"
  />
</div>

<style lang="postcss">
  .wrapper:has(input:focus) {
    @apply border-primary;
  }
</style>
