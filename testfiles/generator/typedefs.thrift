namespace go gentest

typedef binary Binary
typedef string String
typedef i32    Int32

struct St {
	1: Binary b (go.tag="json:\"b_field\" xml:\"BField\""), 
	2: optional String S,
	3: Int32 i
}
