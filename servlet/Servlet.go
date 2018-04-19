package servlet

type Servlet interface {

	/**
	 * Init is called by the servlet container to indicate to a servlet that the servlet
	 * is being placed into service.
	 */
	Init(config ServletConfig)

	/**
	 *
	 * GetServletConfig returns a ServletConfig object, which contains initialization and
	 * startup parameters for this servlet. The ServletConfig
	 * object returned is the one passed to the init method.
	 */
	GetServletConfig() ServletConfig

	/**
	 * Service is called by the servlet container to allow the servlet to respond to a
	 * request.
	 */
	Service(req ServletRequest, res ServletResponse)

	/**
	 * GetServletInfo returns information about the servlet, such as author, version, and
	 * copyright.
	 */
	GetServletInfo() string

	/**
	 * Destroy is called by the servlet container to indicate to a servlet that the servlet
	 * is being taken out of service. This method is only called once all
	 * threads within the servlet's <code>service</code> method have exited or
	 * after a timeout period has passed. After the servlet container calls this
	 * method, it will not call the <code>service</code> method again on this
	 * servlet.
	 */
	Destroy()
}
