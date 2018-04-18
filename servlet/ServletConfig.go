package servlet
/**
 * ServletConfig is a servlet configuration object, used by a servlet container to pass
 * information to a servlet during initialization.
 */
type ServletConfig interface{
	 /**
     * GetServletName returns the name of this servlet instance. The name may be provided via
     * server administration, assigned in the web application deployment
     * descriptor, or for an unregistered (and thus unnamed) servlet instance it
     * will be the servlet's class name.
     */
     GetServletName() string;

	 /**
	  * GetServletContext returns a reference to the ServletContext in which the caller is executing.
	  */
	 GetServletContext() ServletContext;
 
	 /**
	  * GetInitParameter returns a string containing the value of the named
	  * initialization parameter, or nil if the parameter does not
	  * exist.
	  */
	 GetInitParameter(name string) string;
 
	 /**
	  * GetInitParameterNames returns the names of the servlet's initialization parameters as a slice of string objects, or nil 
	  * if the servlet has no initialization parameters.
	  */
	 GetInitParameterNames() []string;

}