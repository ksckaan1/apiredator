<script lang="ts">
  export let requestMethod = "GET";
  export let requestURL = "";

  let requestMethods = [
    {
      value: "GET",
      title: "GET",
      color: "text-green-400",
    },
    {
      value: "POST",
      title: "POST",
      color: "text-amber-400",
    },
    {
      value: "PUT",
      title: "PUT",
      color: "text-sky-400",
    },
    {
      value: "PATCH",
      title: "PATCH",
      color: "text-purple-400",
    },
    {
      value: "DELETE",
      title: "DELETE",
      color: "text-red-400",
    },
    {
      value: "HEAD",
      title: "HEAD",
      color: "text-pink-400",
    },
    {
      value: "TRACE",
      title: "TRACE",
      color: "text-teal-400",
    },
    {
      value: "OPTIONS",
      title: "OPTIONS",
      color: "text-emerald-400",
    },
  ];
</script>

<div
  class="border border-white/20 rounded bg-accent-bg px-3 py-2 h-18 w-full flex gap-3"
>
  <div class={requestMethods.find((r) => r.value === requestMethod)?.color}>
    {requestMethod}
  </div>
  <div>{requestURL}</div>
</div>
