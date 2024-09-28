<script lang="ts">
  import Toast from "$components/Toast.svelte";
import "$lib/style/app.css";
    import type { Snippet } from "svelte";

  interface Props {
    children?: Snippet;
  }

  let { children }: Props = $props();

  if (!('go' in window)) location.replace('/')
</script>

{@render children?.()}

<Toast/>
