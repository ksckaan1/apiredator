<script lang="ts">
	interface Props {
		value?: number;
		label?: string;
		max?: number | null;
		min?: number | null;
		centered?: boolean;
		onPressedEnter?: () => void;
	}

	let {
		value = $bindable(0),
		label = "",
		max = null,
		min = null,
		centered = false,
		onPressedEnter,
	}: Props = $props();

	let inputElem: HTMLInputElement;

	let isValid = $state(true);

	const onInput = () => {
		isValid = false;

		if (!/^[\-]?[0-9]+$/gim.test(value.toString())) return;
		if (max != null && value > max) return;
		if (min != null && value < min) return;

		isValid = true;
	};
</script>

<button
	class="wrapper"
	class:invalid={!isValid}
	onclick={() => inputElem.focus()}
>
	{#if label}
		<span
			class="flex items-center justify-center mr-3 text-white/60 text-nowrap cursor-text"
		>
			{label}
		</span>
	{/if}
	<input
		class:text-center={centered}
		class="w-full bg-transparent focus:outline-none"
		type="text"
		autocorrect="off"
		autocapitalize="none"
		bind:this={inputElem}
		bind:value
		class:invalid={!isValid}
		oninput={onInput}
		onkeypress={(e) => {
			if (e.key === "Enter") onPressedEnter?.();
		}}
	/>
</button>

<style lang="postcss">
	@reference "$styles/app.css";

	.wrapper {
		@apply flex items-center h-10 px-3 border rounded bg-accent-bg border-white/20 cursor-text;
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
