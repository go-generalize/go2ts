export type Data = {
	A: number;
	Array: number[] | null;
	C: string;
	D: number | null;
	Foo?: {
		V: number;
	};
	Map: {[key: string]: "Failure" | "OK"};
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
