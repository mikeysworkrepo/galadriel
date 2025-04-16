export namespace backend {
	
	export class Computer {
	    Name: string;
	    IP: string;
	    Status: string;
	
	    static createFrom(source: any = {}) {
	        return new Computer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.IP = source["IP"];
	        this.Status = source["Status"];
	    }
	}

}

