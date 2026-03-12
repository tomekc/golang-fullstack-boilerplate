<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        children: Snippet;
    }

    let { children }: Props = $props();
</script>

<aside class="aside is-placed-left is-expanded">
    {@render children()}
</aside>
