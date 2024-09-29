<script lang="ts">
  import Section from "$components/ui/Section.svelte";
  import NumberInput from "$components/ui/NumberInput.svelte";
  import SegmentedSelect from "$components/ui/SegmentedSelect.svelte";
  import DurationSelector from "$components/ui/DurationSelector.svelte";

  interface Props {
    activeTestType?: string;
    numberOfRequestsValue?: number;
    numberOfClientsValue?: number;
    testDuration?: string;
    requestTimeout?: string;
    keepAlive?: boolean;
  }

  let {
    activeTestType = $bindable("count"),
    numberOfRequestsValue = $bindable(1),
    numberOfClientsValue = $bindable(1),
    testDuration = $bindable(""),
    requestTimeout = $bindable(""),
    keepAlive = $bindable(true),
  }: Props = $props();

  let testTypes = [
    {
      title: "Count",
      value: "count",
      icon: "mdi:counter"
    },
    {
      title: "Duration",
      value: "duration",
      icon: "gg:sand-clock"
    },
  ];
</script>

<div class="grid grid-cols-1 md:grid-cols-2 w-full mt-5 gap-5">
  <Section title="Test Type" subtitle="Which test type you are using?">
    {#snippet tail()}
      <div class="w-min">
        <SegmentedSelect bind:activeItem={activeTestType} items={testTypes} />
      </div>
    {/snippet}
  </Section>
  <Section
    title="Keep Alive"
    subtitle="Keep-alive reduces latency and network overhead by reusing existing connections."
    isActive={keepAlive}
  ></Section>
  <div>
    {#if activeTestType === "count"}
        <Section
          title="Number of Requests"
          subtitle="How many requests per client?"
        >
          <NumberInput
            bind:value={numberOfRequestsValue}
            min={1}
            max={100000000000}
          />
        </Section>
    {:else if activeTestType === "duration"}
        <Section
          title="Duration"
          subtitle="How long will requests be sent in total?"
        >
          <DurationSelector
            bind:value={testDuration}
            placeholder="Required"
            required
          />
          <div class="mt-3 text-white/50 text-sm">
            <b>Examples:</b> 1h45m, 5m200ms, 1m, 30s <br />
            <b>Allowed Units:</b> h (hour), m (minute), s (second), ms (millisecond),
            us or µs (microsecond), ns (nanosecond)
          </div>
        </Section>
    {/if}
  </div>
  <Section
    title="Number of Clients"
    subtitle="How many clients will make requests at the same time?"
  >
    <NumberInput bind:value={numberOfClientsValue} min={1} max={100000000000} />
  </Section>
  <Section
    title="Request Timeout"
    subtitle="What is the maximum timeout for each request?"
  >
    <DurationSelector
      bind:value={requestTimeout}
      placeholder="Leave blank for no timeout"
    />
    <div class="mt-3 text-white/50 text-sm">
      <b>Examples:</b> 1h45m, 5m200ms, 1m, 30s <br />
      <b>Allowed Units:</b> h (hour), m (minute), s (second), ms (millisecond), us
      or µs (microsecond), ns (nanosecond)
    </div>
  </Section>
</div>
