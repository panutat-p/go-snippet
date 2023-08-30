package object

// https://blog.stackademic.com/writing-a-partitioned-cache-using-go-map-x3-faster-than-the-standard-map-dbfe704fe4bf
type PartitionedMap struct {
   partsnum   uint           // partitions number
   partitions []*partition   // partitions slice
   finder     partitioner    // abstract partitions finder
}

type partition struct {
   stor map[string]any
   sync.RWMutex
}
