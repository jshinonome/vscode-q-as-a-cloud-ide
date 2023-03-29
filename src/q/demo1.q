// go to definition
.file.find

// signature help
.file.find

// rename - pretty bad example p1 -> player1, s1 -> score, w -> window
// semantic highlight
{[b;p1;p2;s1;s2]
    if[~&/(s1;s2)<21;:(s1;s2)>20];
    k:21/:(b;p1;p2;s1;s2);
    if[k in !c;:c[k]];
    w:$[b;
        +/(. a)*.z.s[~b;;p2;;s2]'[p;s1+p:@[;&p>10;-[;10]]p:p1+(!a)];
        +/(. a)*.z.s[~b;p1;;s1;]'[p;s2+p:@[;&p>10;-[;10]]p:p2+(!a)]
    ];
    if[#w;c[k]:w];
    w
 }
