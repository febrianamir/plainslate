export namespace dto {
	
	export class SearchInFilesReq {
	    query: string;
	    is_case_sensitive: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SearchInFilesReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.query = source["query"];
	        this.is_case_sensitive = source["is_case_sensitive"];
	    }
	}

}

export namespace lib {
	
	export class Config {
	    UserHomeDir: string;
	    root_path: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.UserHomeDir = source["UserHomeDir"];
	        this.root_path = source["root_path"];
	    }
	}

}

export namespace model {
	
	export class Match {
	    file_path: string;
	    matches: string[];
	
	    static createFrom(source: any = {}) {
	        return new Match(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.file_path = source["file_path"];
	        this.matches = source["matches"];
	    }
	}
	export class Node {
	    name: string;
	    path: string;
	    type: string;
	    children?: Node[];
	    is_root: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Node(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.path = source["path"];
	        this.type = source["type"];
	        this.children = this.convertValues(source["children"], Node);
	        this.is_root = source["is_root"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

