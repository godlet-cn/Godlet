package servlet

type ServletContext interface{
	 /**
     * RegisterServlet registers Servlet at complie time
     */
	RegisterServlet(servlet Servlet)

}