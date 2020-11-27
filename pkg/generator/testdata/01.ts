export type Data = {
	A: number;
	Array: number[] | null;
	C: string;
	D: number | null;
	EnumArray: ("a" | "b" | "c")[];
	Foo?: {
		V: number;
	};
	Map: {[key: string]: "Failure" | "OK"};
	OptionalArray: (string | null)[];
	Package: {
		data: number;
	} | null;
	Status: "Failure" | "OK";
	Time: string;
	U: {
		Data: number;
	};
	b?: number;
	foo?: number;
}
export type Embedded = {
	foo?: number;
}
