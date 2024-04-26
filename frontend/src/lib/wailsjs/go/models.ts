export namespace domain {
	
	export class KeyValueData {
	    is_active: boolean;
	    key: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new KeyValueData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.is_active = source["is_active"];
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	}
	export class FormData {
	    is_active: boolean;
	    key: string;
	    text_value: string;
	    file_value: string[];
	    row_type: string;
	
	    static createFrom(source: any = {}) {
	        return new FormData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.is_active = source["is_active"];
	        this.key = source["key"];
	        this.text_value = source["text_value"];
	        this.file_value = source["file_value"];
	        this.row_type = source["row_type"];
	    }
	}
	export class Body {
	    type: string;
	    language: string;
	    raw_value: string;
	    binary: string[];
	    formdata: FormData[];
	    xwwwformdata: KeyValueData[];
	
	    static createFrom(source: any = {}) {
	        return new Body(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.language = source["language"];
	        this.raw_value = source["raw_value"];
	        this.binary = source["binary"];
	        this.formdata = this.convertValues(source["formdata"], FormData);
	        this.xwwwformdata = this.convertValues(source["xwwwformdata"], KeyValueData);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Duration {
	    is_duration_active: boolean;
	    hours: number;
	    minutes: number;
	    seconds: number;
	
	    static createFrom(source: any = {}) {
	        return new Duration(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.is_duration_active = source["is_duration_active"];
	        this.hours = source["hours"];
	        this.minutes = source["minutes"];
	        this.seconds = source["seconds"];
	    }
	}
	export class Options {
	    duration: Duration;
	    number_of_clients: number;
	    number_of_requests: number;
	
	    static createFrom(source: any = {}) {
	        return new Options(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.duration = this.convertValues(source["duration"], Duration);
	        this.number_of_clients = source["number_of_clients"];
	        this.number_of_requests = source["number_of_requests"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Request {
	    method: string;
	    url: string;
	    header: KeyValueData[];
	    body: Body;
	
	    static createFrom(source: any = {}) {
	        return new Request(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.method = source["method"];
	        this.url = source["url"];
	        this.header = this.convertValues(source["header"], KeyValueData);
	        this.body = this.convertValues(source["body"], Body);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Data {
	    request: Request;
	    options: Options;
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.request = this.convertValues(source["request"], Request);
	        this.options = this.convertValues(source["options"], Options);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	
	
	
	export class Stat {
	    completed: number;
	    request_per_second: number;
	    status_codes: {[key: number]: number};
	    // Go type: time
	    started_at: any;
	    // Go type: time
	    ended_at: any;
	    duration: number;
	
	    static createFrom(source: any = {}) {
	        return new Stat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.completed = source["completed"];
	        this.request_per_second = source["request_per_second"];
	        this.status_codes = source["status_codes"];
	        this.started_at = this.convertValues(source["started_at"], null);
	        this.ended_at = this.convertValues(source["ended_at"], null);
	        this.duration = source["duration"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

