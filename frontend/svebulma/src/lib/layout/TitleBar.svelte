<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        breadcrumbs?: string[];
        actions?: Snippet;
    }

    let { breadcrumbs = [], actions }: Props = $props();
</script>

<section class="section is-title-bar">
    <div class="level">
        <div class="level-left">
            <div class="level-item">
                <ul>
                    {#each breadcrumbs as crumb}
                        <li>{crumb}</li>
                    {/each}
                </ul>
            </div>
        </div>
        {#if actions}
            <div class="level-right">
                <div class="level-item">
                    <div class="buttons is-right">
                        {@render actions()}
                    </div>
                </div>
            </div>
        {/if}
    </div>
</section>
