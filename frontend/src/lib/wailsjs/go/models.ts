export namespace models {
	
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
	export class RPS {
	    list: number[];
	    latest: number;
	    min: number;
	    avg: number;
	    max: number;
	
	    static createFrom(source: any = {}) {
	        return new RPS(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.list = source["list"];
	        this.latest = source["latest"];
	        this.min = source["min"];
	        this.avg = source["avg"];
	        this.max = source["max"];
	    }
	}
	export class Stat {
	    sent_count: number;
	    rps: RPS;
	    status_codes: {[key: number]: number};
	    // Go type: time
	    started_at: any;
	    // Go type: time
	    ended_at: any;
	    passed_duration: string;
	
	    static createFrom(source: any = {}) {
	        return new Stat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.sent_count = source["sent_count"];
	        this.rps = this.convertValues(source["rps"], RPS);
	        this.status_codes = source["status_codes"];
	        this.started_at = this.convertValues(source["started_at"], null);
	        this.ended_at = this.convertValues(source["ended_at"], null);
	        this.passed_duration = source["passed_duration"];
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
	export class Options {
	    test_type: string;
	    test_duration: string;
	    request_timeout: string;
	    number_of_clients: number;
	    number_of_requests: number;
	    keep_alive: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Options(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.test_type = source["test_type"];
	        this.test_duration = source["test_duration"];
	        this.request_timeout = source["request_timeout"];
	        this.number_of_clients = source["number_of_clients"];
	        this.number_of_requests = source["number_of_requests"];
	        this.keep_alive = source["keep_alive"];
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
	export class Bookmark {
	    id: string;
	    // Go type: time
	    create_at: any;
	    title: string;
	    request: Request;
	    options: Options;
	    stat: Stat;
	    tags: string[];
	
	    static createFrom(source: any = {}) {
	        return new Bookmark(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.create_at = this.convertValues(source["create_at"], null);
	        this.title = source["title"];
	        this.request = this.convertValues(source["request"], Request);
	        this.options = this.convertValues(source["options"], Options);
	        this.stat = this.convertValues(source["stat"], Stat);
	        this.tags = source["tags"];
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
	export class BookmarkList {
	    bookmarks: Bookmark[];
	    limit: number;
	    offset: number;
	    count: number;
	
	    static createFrom(source: any = {}) {
	        return new BookmarkList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.bookmarks = this.convertValues(source["bookmarks"], Bookmark);
	        this.limit = source["limit"];
	        this.offset = source["offset"];
	        this.count = source["count"];
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
	
	
	
	
	

}

