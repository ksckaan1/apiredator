<script lang="ts">
  export let value = 0;
  export let label = "";
  export let max: number | null = null;
  export let min: number | null = null;

  let inputElem: HTMLInputElement;

  let isValid = true;

  const onInput = () => {
    isValid = false;

    if (!/^[\-]?[0-9]+$/gim.test(value.toString())) return;
    if (max != null && value > max) return;
    if (min != null && value < min) return;

    isValid = true;
  };
</script>

<div class="wrapper" class:invalid={!isValid}>
  {#if label}
    <button
      class="flex items-center justify-center mr-3 text-white/60 text-nowrap cursor-text"
      on:click={() => inputElem.focus()}
    >
      {label}
    </button>
  {/if}
  <input
    class="w-full bg-transparent focus:outline-none"
    type="text"
    autocorrect="off"
    autocapitalize="none"
    bind:this={inputElem}
    bind:value
    class:invalid={!isValid}
    on:input={onInput}
  />
</div>

<style lang="postcss">
  .wrapper {
    @apply flex items-center h-10 px-3 border rounded bg-default-bg border-white/20;
  }

  .wrapper:has(input:focus) {
    @apply border-primary;
  }

  .wrapper.invalid {
    @apply border-red-500;
  }
  input.invalid {
    @apply text-red-500;
  }
</style>
