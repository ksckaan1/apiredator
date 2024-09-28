<script lang="ts">
  interface Props {
    isActive?: boolean;
    onclick?: (e: MouseEvent) => void;
  }

  let { isActive = false, onclick }: Props = $props();
</script>

<button {onclick}>
  {#if isActive}
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      width="2rem"
      height="2rem"><path fill="currentColor" d="M5 21V3h14v18l-7-3z" /></svg
    >
  {:else}
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      width="2rem"
      height="2rem"
      ><path
        fill="currentColor"
        d="M5 21V5q0-.825.588-1.412T7 3h10q.825 0 1.413.588T19 5v16l-7-3zm2-3.05l5-2.15l5 2.15V5H7zM7 5h10z"
      /></svg
    >
  {/if}
</button>
