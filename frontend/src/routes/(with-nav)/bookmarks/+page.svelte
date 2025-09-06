<script lang="ts">
	import type { models } from "$lib/wailsjs/go/models";
	import { fade, slide } from "svelte/transition";
	import {
		DeleteBookmarks,
		GetAllBookmarks,
		GetAllTags,
	} from "$lib/wailsjs/go/service/AppService";
	import { showToast } from "$stores/toast";
	import DeleteBookmarksModal from "./parts/DeleteBookmarksModal.svelte";
	import { goto } from "$app/navigation";
	import BookmarkCard from "./parts/BookmarkCard.svelte";
	import BookmarkHeader from "./parts/BookmarkHeader.svelte";
	import BookmarkPaginator from "./parts/BookmarkPaginator.svelte";

	let bookmarks: models.Bookmark[] = $state([]);
	let searchValue = $state("");
	let tagValue = $state("");
	let allTags: any[] = $state([]);
	let count = $state(0);
	let limit = $state(10);
	let offset = $state(0);

	let bouncedSearchValue = $state("");
	let selectedBookmarks: string[] = $state([]);

	let showDeleteBookmarkModal = $state(false);

	let timer: number;

	$effect(() => {
		if (searchValue != undefined) {
			clearTimeout(timer);
			timer = setTimeout(() => {
				bouncedSearchValue = searchValue;
			}, 750);
		}
	});

	const makeSearch = (offsetParam: number) => {
		GetAllBookmarks(bouncedSearchValue, tagValue, 10, offsetParam).then(
			(result) => {
				bookmarks = result.bookmarks;
				count = result.count;
				offset = result.offset;
				count = result.count;
			},
		);
	};

	makeSearch(0);

	$effect(() => {
		if (bouncedSearchValue != undefined && tagValue != undefined) {
			makeSearch(0);
		}
	});

	GetAllTags().then((result) => {
		allTags = [
			{ title: "All Tags", value: "" },
			...result.map((t) => ({ value: t, title: t })),
		];
	});

	const deleteBookmarks = () => {
		DeleteBookmarks(selectedBookmarks)
			.then(() => {
				showToast({
					type: "success",
					message: `${selectedBookmarks.length} bookmark(s) deleted`,
				});
				selectedBookmarks = [];
				showDeleteBookmarkModal = false;
				makeSearch(0);
			})
			.catch((e: Error) => {
				showToast({
					type: "error",
					message: e.message,
				});
				selectedBookmarks = [];
				showDeleteBookmarkModal = false;
				makeSearch(0);
			});
	};

	const onClickBookmark = (id: string) => {
		goto(`/inspect?id=${id}`);
	};

	const onBookmarkSelected = (id: string, selected: boolean) => {
		if (selected && !selectedBookmarks.includes(id)) {
			selectedBookmarks.push(id);
			return;
		}

		selectedBookmarks = selectedBookmarks.filter(
			(bookmarkId) => bookmarkId !== id,
		);
	};
</script>

<div
	in:fade={{ duration: 200, delay: 200 }}
	out:fade={{ duration: 200 }}
	class="h-[calc(100vh-40px)] flex flex-col"
>
	<div class="flex-1 hide-scrollbar overflow-y-auto flex flex-col px-5 pt-5">
		<div class="mx-auto w-full max-w-7xl flex flex-col gap-4 h-screen">
			<h1 class="text-3xl">Bookmarks</h1>
			<BookmarkHeader
				bind:searchValue
				bind:tagValue
				{allTags}
				bind:showDeleteBookmarkModal
				bind:selectedBookmarks
			/>
			{#if bookmarks && bookmarks.length > 0}
				<div
					class:pb-5={count < 11}
					class="flex flex-col pr-[1px] gap-4 {count > 10
						? 'h-[calc(100vh-19rem)]'
						: 'h-[calc(100vh-14rem)]'} hide-scrollbar overflow-y-scroll"
				>
					{#each bookmarks as bookmark (bookmark.id)}
						<BookmarkCard
							{bookmark}
							{onClickBookmark}
							bind:tagValue
							selected={selectedBookmarks.includes(bookmark.id)}
							onBookmarkSelected={(selected) =>
								onBookmarkSelected(bookmark.id, selected)}
						/>
					{/each}
				</div>
				{#if count > 10}
					<BookmarkPaginator
						{limit}
						{offset}
						{count}
						onOffsetChange={(offset) => makeSearch(offset)}
					/>
				{/if}
			{:else}
				<div
					class="text-center flex flex-col justify-center items-center text-white/50"
				>
					There is no bookmark added!
				</div>
			{/if}
		</div>
	</div>
</div>

<DeleteBookmarksModal
	bind:showModal={showDeleteBookmarkModal}
	{selectedBookmarks}
	{deleteBookmarks}
/>
