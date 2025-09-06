<script lang="ts">
	import {
		getOnPrimaryColor,
		getPrimaryColor,
		requestMethods,
	} from "$lib/request-methods";
	import type { models } from "$lib/wailsjs/go/models";
	import Icon from "@iconify/svelte";
	interface Props {
		bookmark: models.Bookmark;
		onClickBookmark: (id: string) => void;
		tagValue: string;
		selected: boolean;
		onBookmarkSelected: (selected: boolean) => void;
	}

	let {
		bookmark,
		onClickBookmark,
		tagValue = $bindable(""),
		selected = $bindable(false),
		onBookmarkSelected,
	}: Props = $props();
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	class="flex border overflow-hidden bg-accent-bg border-white/20 rounded flex-shrink-0 hover:bg-white/5 cursor-pointer group"
	onclick={() => onClickBookmark(bookmark.id)}
>
	<div
		style="background-color: {getPrimaryColor(
			requestMethods,
			bookmark.request.method,
		)}"
		class="flex w-8 pl-1 items-center flex-shrink-0"
	>
		<div
			style="color: {getOnPrimaryColor(
				requestMethods,
				bookmark.request.method,
			)}"
			class="[writing-mode:vertical-lr] rotate-180"
		>
			{bookmark.request.method}
		</div>
	</div>
	<div class="p-5 flex-1">
		<div class="text-lg text-primary">
			{bookmark.title || "Untitled Bookmark"}
		</div>
		<div class="flex gap-3">
			<div>
				{bookmark.request.url}
			</div>
		</div>
		<div class="mt-2 text-white/30 font-light">
			{new Date(bookmark.create_at).toString()}
		</div>
		<div class="mt-1 flex flex-wrap gap-3 items-center">
			<Icon icon="mdi:tags" class="text-white/50" />
			{#if bookmark.tags && bookmark.tags.length > 0}
				{#each bookmark.tags as tag}
					<button
						onclick={(e) => {
							e.stopPropagation();
							tagValue = tag;
						}}
						class="text-white/50 transition-colors duration-200 hover:text-primary"
					>
						#{tag}
					</button>
				{/each}
			{:else}
				<div class="text-white/50">No Tag</div>
			{/if}
		</div>
	</div>
	<div class="flex flex-col justify-center pr-3 flex-shrink-0">
		<input
			type="checkbox"
			id={bookmark.id}
			value={bookmark.id}
			checked={selected}
			class="hidden"
			onclick={(e) => {
				e.stopPropagation();
				onBookmarkSelected(!selected);
			}}
		/>
		<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
		<label
			for={bookmark.id}
			class="cursor-pointer"
			onclick={(e) => e.stopPropagation()}
		>
			{#if selected}
				<Icon icon="bxs:checkbox" width="40" class="text-primary" />
			{:else}
				<Icon
					icon="bx:checkbox"
					width="40"
					class="text-white/30 opacity-0 group-hover:opacity-100"
				/>
			{/if}
		</label>
	</div>
</div>
