# BitSet
  Easy bit operation with a single value undertable

## Usage
Init value with 0.
```
bitSet := BitSet{0}
```

Set position 0, 1, 2 as 1. Value is now `b'111'`, equals to `7`.
```
bitSet := BitSet{0}
bitSet.Set(0, 1, 2)
```

Clear position 1, 2 as 0. Value is now `b'001'`, equals to `1`
```
bitSet := BitSet{7}
bitSet.Clear(1, 2)
```

Toggle position 0, 2 as 0. Value is now `b'010'`, equals to `2`
```
bitSet := BitSet{7}
bitSet.Toggle(0, 2)
bitSet.IsSet(0) // return false
bitSet.IsSet(1) // return true
bitSet.IsSet(2) // return false
```

