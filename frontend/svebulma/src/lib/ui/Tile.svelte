<script lang="ts">
    type Color = 'primary' | 'link' | 'info' | 'success' | 'warning' | 'danger';

    interface Props {
        label: string;
        value: string;
        icon: string;
        color?: Color;
    }

    let { label, value, icon, color = 'primary' }: Props = $props();

    const iconColorClass = $derived(`has-text-${color}`);
</script>

<div class="column">
    <div class="card">
        <div class="card-content">
            <div class="level is-mobile">
                <div class="level-item">
                    <div class="is-widget-label">
                        <h3 class="subtitle is-spaced">{label}</h3>
                        <h1 class="title">{value}</h1>
                    </div>
                </div>
                <div class="level-item has-widget-icon">
                    <div class="is-widget-icon">
                        <span class="icon is-large {iconColorClass}">
                            <i class="mdi mdi-{icon} mdi-48px"></i>
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
