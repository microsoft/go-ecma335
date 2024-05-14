package genwinsyscallproto

import (
	"github.com/microsoft/go-winmd"
)

type typeNameKey struct {
	NamespaceStart uint32
	NameStart      uint32
}

func typeRefKey(ref *winmd.TypeRef) typeNameKey {
	return typeNameKey{ref.Namespace.Start, ref.Name.Start}
}

func typeDefKey(def *winmd.TypeDef) typeNameKey {
	return typeNameKey{def.Namespace.Start, def.Name.Start}
}

type typeDefCache struct {
	// resolved maps type name keys -> resolved TypeDef information.
	resolved map[typeNameKey]*resolvedDef
	// unresolved maps type name keys -> winmd.TypeDef index.
	unresolved map[typeNameKey]winmd.Index

	// Duplicated TypeDefs are an uncommon special case, treat them separately to avoid
	// allocating a slice for every TypeDef, only for those that are duplicated.
	// WinMD TypeDefs use duplicated TypeDef names to represent functions with the same name but
	// different signatures due to architecture-specific overloads.

	// resolvedDuplicated maps type name keys -> list of TypeDef indices with the same name.
	resolvedDuplicated map[typeNameKey][]*resolvedDef
	// unresolvedDuplicated maps type name keys -> list of winmd.TypeDef indices.
	unresolvedDuplicated map[typeNameKey][]winmd.Index
}

func newTypeDefCache() *typeDefCache {
	return &typeDefCache{
		resolved:             make(map[typeNameKey]*resolvedDef),
		unresolved:           make(map[typeNameKey]winmd.Index),
		unresolvedDuplicated: make(map[typeNameKey][]winmd.Index),
		resolvedDuplicated:   make(map[typeNameKey][]*resolvedDef),
	}
}

func (tc *typeDefCache) add(i winmd.Index, typ *winmd.TypeDef) {
	key := typeDefKey(typ)
	if _, ok := tc.unresolvedDuplicated[key]; ok {
		tc.unresolvedDuplicated[key] = append(tc.unresolvedDuplicated[key], i)
	} else if _, ok := tc.unresolved[key]; ok {
		tc.unresolvedDuplicated[key] = append(tc.unresolvedDuplicated[key], tc.unresolved[key], i)
		delete(tc.unresolved, key)
	} else {
		tc.unresolved[key] = i
	}
}

func (tc *typeDefCache) resolve(r *resolvedDef) {
	key := typeNameKey{r.Namespace.Start, r.Name.Start}
	if tc.unresolvedDuplicated[key] == nil {
		tc.resolved[key] = r
		delete(tc.unresolved, key)
	} else {
		tc.resolvedDuplicated[key] = append(tc.resolvedDuplicated[key], r)
		delete(tc.unresolved, key)
	}
}

func (tc *typeDefCache) get(namespace, name winmd.String, arch Arch) *resolvedDef {
	key := typeNameKey{namespace.Start, name.Start}
	if defs, ok := tc.resolvedDuplicated[key]; ok {
		for _, def := range defs {
			if def.Arch&arch == arch {
				return def
			}
		}
	}
	if def, ok := tc.resolved[key]; ok {
		return def
	}
	return nil
}

func (tc *typeDefCache) collect(filter func(r *resolvedDef) bool) []*resolvedDef {
	ret := make([]*resolvedDef, 0, len(tc.resolved)+len(tc.resolvedDuplicated))
	for _, r := range tc.resolved {
		if filter(r) {
			ret = append(ret, r)
		}
	}
	for _, def := range tc.resolvedDuplicated {
		for _, r := range def {
			if filter(r) {
				ret = append(ret, r)
			}
		}
	}
	return ret
}
