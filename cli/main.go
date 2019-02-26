package main

import (
	"fmt"
	"os"

	"github.com/keptn/keptn/cli/cmd"
	"github.com/keptn/keptn/cli/utils"
)

func init() {
	utils.Init(os.Stdout, os.Stdout, os.Stderr)
}

const logo = `                                                                                                                                     
                ##########*                                                                                                                                    
           ,#############    ##                                                                                                                                
       (###############    ####    *                                                                                                                           
    ##################    ###*    ###.                                                                                                                         
   #######      ####    ####    ####                                                                                                                           
   #####          ,   (###    ####    ##                 .&&&&                                                                                                 
  (####   #####      ####    ####    ###                 .&&&&                                                                                                 
  #####    ####    ####    ####    ####                  .&&&&                                                              &&&&&                              
 .######         .###    *###    ####                    .&&&&                                                              &&&&&                              
 ##########     ####    ####    ####    #(               .&&&&                                                              &&&&&                              
 #########    ####    ####    ####    ####               .&&&&       &&&&&/       &&&&&&&&&&/        &&&&&&&&&&&&&%         &&&&&&&&&&&&,     &&&&&&&&&&&&&&   
#########    ####    ####   .###/   /#####               .&&&&     &&&&&&       &&&&&&&&&&&&&&%      &&&&&&&&&&&&&&&&       &&&&&&&&&&&&,     &&&&&&&&&&&&&&&& 
#######    ####    ####    ####    ########              .&&&&   &&&&&&        &&&&&.     /&&&&&     &&&&&       &&&&&(     &&&&&             &&&&&      &&&&&&
 ####(   .###    (###    ####    #########               .&&&& &&&&&&         &&&&&        *&&&&     &&&&&        &&&&&     &&&&&             &&&&&       &&&&&
  ##    ####    ####    ####    ########                 .&&&&&&&&&           &&&&&&&&&&&&&&&&&&     &&&&&         &&&&&    &&&&&             &&&&&       &&&&&
      ####    ####    ####    #########                  .&&&&&&&&&&          &&&&&&&&&&&&&&&&&&     &&&&&         &&&&&    &&&&&             &&&&&       &&&&&
     ####    ###/   (###,   (########                    .&&&&  &&&&         &&&&&                  &&&&&         &&&&&    &&&&&             &&&&&       &&&&&
           ####    ####    ########*                     .&&&&   .&&&&&       #&&&&&                 &&&&&        &&&&&     /&&&&             &&&&&       &&&&&
         ####    ####    #########                       .&&&&     &&&&&&      &&&&&&&%    ,&&&      &&&&&&&( .&&&&&&&       &&&&&&/  %&&     &&&&&       &&&&&
          ##    ####    ########                         .&&&&       &&&&      &&&&&&&&&&&&&&      &&&&&&&&&&&&&&&&         &&&&&&&&&&&     &&&&&       &&&&&
                                                                                    .&&&&&&&*        &&&&&  *&&&&                 (&&%                         
                                                                                                     &&&&&                                                     
                                                                                                     &&&&&                                                     
                                                                                                     &&&&&`

func main() {
	fmt.Println(logo)
	cmd.Execute()
}
