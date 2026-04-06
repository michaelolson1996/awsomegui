export namespace awsdto {
	
	export class AWSProfileDTO {
	    name: string;
	    isDefault: boolean;
	    region: string;
	
	    static createFrom(source: any = {}) {
	        return new AWSProfileDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.isDefault = source["isDefault"];
	        this.region = source["region"];
	    }
	}
	export class VPCDTO {
	    id: string;
	    name: string;
	    cidrBlock: string;
	    state: string;
	    isDefault: boolean;
	    region: string;
	
	    static createFrom(source: any = {}) {
	        return new VPCDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.cidrBlock = source["cidrBlock"];
	        this.state = source["state"];
	        this.isDefault = source["isDefault"];
	        this.region = source["region"];
	    }
	}

}

