Exporter
===

exporter exports metrics of Zai to prometheus.

# Metric Lib

The official prometheus client is being used, not like other places which using
[victoria_metrics](https://github.com/VictoriaMetrics/metrics).

For other components, they need high speed and want the block as small as possible. The victoria_metrics is a better choice.

For exporter, we want richer and more flexible labels supporting. I prefer the official lib.

This inconsistency is acceptable in my opinion.

