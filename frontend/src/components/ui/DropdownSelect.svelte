<script lang="ts">
	import { getPrimaryColor, getTitle } from "$lib/request-methods";
	import { slide } from "svelte/transition";

	interface Props {
		value?: string;
		items?: any[];
		transparent?: boolean;
	}

	let {
		value = $bindable(""),
		items = [],
		transparent = false,
	}: Props = $props();

	let elem: HTMLDivElement;
	let isFocused = $state(false);

	const selectItem = (v: string) => {
		value = v;
		elem.blur();
	};
</script>

<div
	class="relative w-full group"
	bind:this={elem}
	tabindex="-1"
	onfocus={() => (isFocused = true)}
	onblur={() => (isFocused = false)}
>
	<div
		class="flex items-center justify-between w-full h-10 px-3 cursor-pointer {!transparent
			? 'bg-accent-bg hover:bg-white/5 border rounded border-white/20'
			: ''}"
	>
		<span style="color: {getPrimaryColor(items, value)}">
			{getTitle(items, value) ?? value.toUpperCase()}
		</span>
		<i
			class="transition-all duration-200 fa-solid fa-chevron-down group-focus:rotate-180"
		></i>
	</div>
	{#if isFocused}
		<div
			transition:slide={{ duration: 200 }}
			class="absolute top-0 left-0 z-50 w-full border rounded bg-accent-bg border-white/20 max-h-96 overflow-y-scroll hide-scrollbar"
		>
			<button
				style="color: {getPrimaryColor(items, value)}"
				class="block w-full px-3 py-2 text-left hover:bg-white/5"
				onclick={() => selectItem(value)}
			>
				{getTitle(items, value) ?? value.toUpperCase()}
			</button>
			{#each items as item}
				{#if item.value != value}
					<button
						style="color: {getPrimaryColor(items, item.value)}"
						class="block w-full px-3 py-2 text-left hover:bg-white/5"
						onclick={() => selectItem(item.value)}
					>
						{item.title}
					</button>
				{/if}
			{/each}
		</div>
	{/if}
</div>
