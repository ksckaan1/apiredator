<script lang="ts">
  import { slide, fade } from "svelte/transition";
  import { goto } from "$app/navigation";
  import logo from "$assets/logo.svg";
  import { onMount } from "svelte";
  onMount(() => {
    setTimeout(() => {
      goto("/new-request", {
        replaceState: true,
      });
    }, 2000);
  });
</script>

<div
  style="--wails-draggable:drag"
  class="w-screen h-screen flex justify-center items-center relative overflow-hidden"
  in:fade={{ delay: 200, duration: 200 }}
  out:fade={{ duration: 200 }}
>
  <img
    in:fade={{ delay: 200, duration: 200 }}
    out:slide={{ duration: 200 }}
    src={logo}
    alt="logo"
    class="w-48"
  />
  <div class="absolute top-0 left-0 w-screen h-screen overflow-hidden"></div>
</div>
