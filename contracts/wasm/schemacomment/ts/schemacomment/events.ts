// COPYRIGHT OF A TEST SCHEMA DEFINITION 1
// COPYRIGHT OF A TEST SCHEMA DEFINITION 2

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as wasmtypes from "wasmlib/wasmtypes";

export class SchemaCommentEvents {

	// header comment for TestEvent 1
// header comment for TestEvent 2
	testEvent(
	// header comment for eventParam1 1
	// header comment for eventParam1 2
		eventParam1: string,
	// header comment for eventParam2 1
	// header comment for eventParam2 2
		eventParam2: string,
	): void {
		const evt = new wasmlib.EventEncoder("schemacomment.testEvent");
		evt.encode(wasmtypes.stringToString(eventParam1));
		evt.encode(wasmtypes.stringToString(eventParam2));
		evt.emit();
	}

	// header comment for TestEventNoParams 1
// header comment for TestEventNoParams 2
	testEventNoParams(): void {
		const evt = new wasmlib.EventEncoder("schemacomment.testEventNoParams");
		evt.emit();
	}
}
