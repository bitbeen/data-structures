package data_structures

import "testing"
import (
    "github.com/timtadh/data-structures/types"
    "github.com/timtadh/data-structures/tree"
    "github.com/timtadh/data-structures/hashtable"
)

func TestAvlTreeCast(t *testing.T) {
    tree := tree.NewAvlTree()
    _ = types.Map(tree)
}

func TestHashtableCast(t *testing.T) {
    hash := hashtable.NewHashTable(16)
    _ = types.Sized(hash)
    _ = types.MapIterable(hash)
    _ = types.MapOperable(hash)
    _ = types.Map(hash)
}

func TestLinearHashtableCast(t *testing.T) {
    hash := hashtable.NewLinearHash()
    _ = types.Sized(hash)
    _ = types.MapIterable(hash)
    _ = types.MapOperable(hash)
    _ = types.Map(hash)
}


