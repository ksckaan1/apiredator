<script lang="ts">
	import { getPrimaryColor, requestMethods } from "$lib/request-methods";

	interface Props {
		requestMethod?: string;
		requestURL?: string;
	}

	let {
		requestMethod = $bindable("GET"),
		requestURL = $bindable(""),
	}: Props = $props();
</script>

<div
	class="border border-white/20 rounded bg-accent-bg px-3 py-2 w-full flex gap-3"
>
	<div style="color: {getPrimaryColor(requestMethods, requestMethod)}">
		{requestMethod}
	</div>
	<div>{requestURL}</div>
</div>
