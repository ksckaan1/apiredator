<script lang="ts">
	import Button from "$components/ui/Button.svelte";
	import NumberInput from "$components/ui/NumberInput.svelte";

	interface Props {
		limit: number;
		offset: number;
		count: number;
		onOffsetChange: (offset: number) => void;
	}

	let { limit, offset, count, onOffsetChange }: Props = $props();

	let lastPage = $derived(Math.ceil(count / limit));
	let shownPage = $derived(Math.ceil(offset / limit) + 1);
	let currentPage = $derived(Math.ceil(offset / limit) + 1);

	const onPressedEnter = () => {
		if (shownPage < 1 || shownPage > lastPage) return;
		onOffsetChange((shownPage - 1) * limit);
	};

	const onFirstButtonClicked = () => {
		shownPage = 1;
		onPressedEnter();
	};

	const onPreviousButtonClicked = () => {
		shownPage = Number(shownPage) - 1;
		onPressedEnter();
	};

	const onNextButtonClicked = () => {
		shownPage = Number(shownPage) + 1;
		onPressedEnter();
	};

	const onLastButtonClicked = () => {
		shownPage = lastPage;
		onPressedEnter();
	};
</script>

<div class="h-10 flex items-center justify-center gap-x-2">
	{#if currentPage > 1}
		<Button onclick={onFirstButtonClicked} variant="secondary">First</Button
		>
	{/if}
	{#if currentPage > 1}
		<Button onclick={onPreviousButtonClicked} variant="secondary"
			>Previous</Button
		>
	{/if}
	<div class="w-14">
		<NumberInput
			min={1}
			max={lastPage}
			bind:value={shownPage}
			centered
			{onPressedEnter}
		/>
	</div>
	<span>/ {lastPage}</span>
	{#if currentPage < lastPage}
		<Button onclick={onNextButtonClicked} variant="secondary">Next</Button>
	{/if}
	{#if currentPage < lastPage}
		<Button onclick={onLastButtonClicked} variant="secondary">Last</Button>
	{/if}
</div>
