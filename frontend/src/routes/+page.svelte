<script lang="ts">
  import { goto } from "$app/navigation";
  import logo from "$assets/logo.svg";
  import { onMount } from "svelte";
  onMount(() => {
    setTimeout(() => {
      goto("/new-request", {
        replaceState: true,
      });
    }, 3000);
  });
</script>

<div
  style="--wails-draggable:drag"
  class="w-screen h-screen flex justify-center items-center"
>
  <img src={logo} alt="logo" class="w-48" />
</div>
